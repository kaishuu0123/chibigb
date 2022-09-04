package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strings"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/kaishuu0123/chibigb/chibigb"
	"github.com/kaishuu0123/chibigb/internal/gui"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/draw"
)

const (
	SCALE               int     = 2
	WINDOW_WIDTH        int     = chibigb.GAMEBOY_WIDTH * SCALE
	WINDOW_HEIGHT       int     = chibigb.GAMEBOY_HEIGHT * SCALE
	AUDIO_MASTER_VOLUME float64 = 0.3 // min: 0.0 max: 1.0
)

var (
	isRunning                    = false
	console     *chibigb.Console = nil
	audioDevice sdl.AudioDeviceID
	audioBuffer [735 * 4]byte // *2 for stereo, *2 for sizeof(byte)
)

func main() {
	flag.Parse()
	if len(flag.Args()) >= 1 {
		_, err := os.Stat(flag.Arg(0))
		if err != nil {
			log.Fatalln("No ROM file specified or found")
		}

		ResetConsole(flag.Arg(0))
	}
	defer StopAudio()

	window := gui.NewMasterWindow("ChibiGB", WINDOW_WIDTH, WINDOW_HEIGHT, -1)
	window.SetDropCallback(onDrop)
	screenImage := image.NewRGBA(image.Rect(0, 0, chibigb.GAMEBOY_WIDTH, chibigb.GAMEBOY_HEIGHT))
	displayImage := image.NewRGBA(image.Rect(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT))

	var texture imgui.TextureID
	for !window.Platform.ShouldStop() {
		window.Platform.ProcessEvents()

		if isRunning {
			processInputController(window.Platform.Window, console)

			console.RunToVBlank()
			console.SetPixels(screenImage.Pix)
			draw.NearestNeighbor.Scale(displayImage, displayImage.Bounds(), screenImage, screenImage.Bounds(), draw.Over, nil)

		}

		texture, _ = window.Renderer.CreateImageTexture(displayImage)
		renderGUI(window, &texture)
		window.Renderer.ReleaseImage(texture)

		if isRunning {
			PlayAudio(console)
		}
	}
}

func onDrop(names []string) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s", names[0]))
	dropInFiles := sb.String()
	ResetConsole(dropInFiles)
}

func renderGUI(w *gui.MasterWindow, texture *imgui.TextureID) {
	w.Platform.NewFrame()
	imgui.NewFrame()

	if isRunning {
		imgui.BackgroundDrawList().
			AddImage(
				*texture,
				imgui.Vec2{X: 0, Y: 0},
				imgui.Vec2{X: float32(WINDOW_WIDTH), Y: float32(WINDOW_HEIGHT)},
			)
	} else {
		var msg string = "ChibiGB is \ncurrently stopped.\n\nPlease drag and drop \nROM file."
		textSize := imgui.CalcTextSize(msg, false, 0)
		xpos := (float32(WINDOW_WIDTH) - textSize.X) / 2
		ypos := (float32(WINDOW_HEIGHT) - textSize.Y) / 2
		imgui.ForegroundDrawList().
			AddText(
				imgui.Vec2{X: xpos, Y: ypos},
				imgui.PackedColor(0xFFFFFFFF),
				msg,
			)
	}

	imgui.Render()

	w.Renderer.PreRender(w.ClearColor)
	w.Renderer.Render(w.Platform.DisplaySize(), w.Platform.FramebufferSize(), imgui.RenderedDrawData())
	w.Platform.PostRender()
}

func processInputController(window *glfw.Window, console *chibigb.Console) {
	var result [8]bool
	result[chibigb.KEY_A] = window.GetKey(glfw.KeyZ) == glfw.Press
	result[chibigb.KEY_B] = window.GetKey(glfw.KeyX) == glfw.Press
	result[chibigb.KEY_START] = window.GetKey(glfw.KeyEnter) == glfw.Press
	result[chibigb.KEY_SELECT] = window.GetKey(glfw.KeyRightShift) == glfw.Press
	result[chibigb.KEY_RIGHT] = window.GetKey(glfw.KeyRight) == glfw.Press
	result[chibigb.KEY_LEFT] = window.GetKey(glfw.KeyLeft) == glfw.Press
	result[chibigb.KEY_UP] = window.GetKey(glfw.KeyUp) == glfw.Press
	result[chibigb.KEY_DOWN] = window.GetKey(glfw.KeyDown) == glfw.Press

	for i := 0; i < len(result); i++ {
		console.SetButtonState(i, result[i])
	}
}

func ResetConsole(filePath string) {
	StopAudio()
	isRunning = false

	log.Println("Reset Console")
	log.Printf("ROM file path: %s\n", filePath)
	data, err := readFile(filePath)
	if err != nil {
		log.Fatalf("readFile error: %s\n", err)
	}
	log.Printf("ROM Loaded %s\n", filePath)

	console = chibigb.NewConsole()
	console.LoadFromBuffer(data)
	isRunning = true

	StartAudio()
}

func PlayAudio(console *chibigb.Console) {
	console.SetSoundBuffer(audioBuffer[:])
	if sdl.GetQueuedAudioSize(audioDevice) <= uint32(len(audioBuffer)*6) {
		src := (*[len(audioBuffer) * 2]uint8)(unsafe.Pointer(&audioBuffer[0]))
		dst := make([]uint8, len(audioBuffer)*2)
		volume := int(math.Floor(float64(sdl.MIX_MAXVOLUME) * AUDIO_MASTER_VOLUME))

		sdl.MixAudioFormat(&dst[0], &src[0], sdl.AUDIO_S16, uint32(len(audioBuffer)*2), volume)

		// don't queue audio if buffer is still filled
		sdl.QueueAudio(audioDevice, dst[:len(audioBuffer)])
	}
}

func StartAudio() {
	err := sdl.InitSubSystem(sdl.INIT_AUDIO)
	if err != nil {
		log.Fatalf("Failed to init SDL: %s\n", err)
	}

	var want, have sdl.AudioSpec
	want.Freq = 44100
	want.Format = sdl.AUDIO_S16
	want.Channels = 2
	want.Samples = 1024
	want.Callback = nil // use queue
	audioDevice, err = sdl.OpenAudioDevice("", false, &want, &have, 0)
	if err != nil {
		log.Fatalf("SDL OpenAudioDevice error: %s\n", err)
	}
	sdl.PauseAudioDevice(audioDevice, false)
}

func StopAudio() {
	if isRunning {
		sdl.QuitSubSystem(sdl.INIT_AUDIO)
	}
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())
	if err := binary.Read(file, binary.LittleEndian, data); err != nil {
		return nil, err
	}

	return data, nil
}
