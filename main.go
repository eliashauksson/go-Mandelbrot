package main

import (
	"image/color"
	"log"
	"math"
	"math/cmplx"
	"runtime"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 600
	screenHeight = 400
)

var (
	wg sync.WaitGroup

	img    *ebiten.Image
	imgPix []byte

	preRect     *ebiten.Image
	showPreRect bool

	reStart = -2.0
	reEnd   = 1.0
	imStart = -1.0
	imEnd   = 1.0

	maxIter = 40.0
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		newX, newY := ebiten.CursorPosition()

		newReStart := (float64(newX)/float64(screenWidth)*math.Abs(reStart-reEnd) + reStart) - (0.25 * math.Abs(reStart-reEnd))
		newReEnd := (float64(newX)/float64(screenWidth)*math.Abs(reStart-reEnd) + reStart) + (0.25 * math.Abs(reStart-reEnd))
		newImStart := (float64(newY)/float64(screenHeight)*math.Abs(imStart-imEnd) + imStart) - (0.25 * math.Abs(imStart-imEnd))
		newImEnd := (float64(newY)/float64(screenHeight)*math.Abs(imStart-imEnd) + imStart) + (0.25 * math.Abs(imStart-imEnd))

		reStart = newReStart
		reEnd = newReEnd
		imStart = newImStart
		imEnd = newImEnd

		updateImg()
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		reStart = -2.0
		reEnd = 1.0
		imStart = -1.0
		imEnd = 1.0

		maxIter = 40

		updateImg()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		showPreRect = !showPreRect
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)

	if showPreRect {
		cursX, cursY := ebiten.CursorPosition()
		cursX -= screenWidth / 4
		cursY -= screenHeight / 4

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(cursX), float64(cursY))

		screen.DrawImage(preRect, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func computeMandelbrot() {

	start := time.Now()
	log.Println("++++++++++++++++++++++++++++++++++++")
	log.Println("Computing Mandelbrot...")

	maxIter = maxIter * 1.3

	for x := 0; x < screenWidth; x++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()

			for y := 0; y < screenHeight; y++ {
				var c complex128 = complex(
					reStart+(float64(x)/float64(screenWidth))*(reEnd-reStart),
					imStart+(float64(y)/float64(screenHeight))*(imEnd-imStart),
				)

				var q float64
				var z complex128 = 0

				for i := 0; i < int(maxIter); i++ {
					z = cmplx.Pow(z, 2) + c

					if cmplx.Abs(z) > 2.0 {
						q = float64(i) / float64(maxIter)
						break
					} else {
						q = 0.0
					}
				}
				p := 4 * (x + y*screenWidth)
				col := byte(q * 255)

				if q > 0.5 {
					imgPix[p] = col
					imgPix[p+1] = 255
					imgPix[p+2] = col
					imgPix[p+3] = 255
				} else {
					imgPix[p] = 0
					imgPix[p+1] = col
					imgPix[p+2] = 0
					imgPix[p+3] = 255
				}
			}
		}(x)
	}
	wg.Wait()

	log.Println("Calculation Time:", time.Since(start))
	start = time.Now()

	img.ReplacePixels(imgPix)

	log.Println("Render Time:", time.Since(start))
}

func updateImg() {
	img, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	imgPix = make([]byte, screenWidth*screenHeight*4)

	computeMandelbrot()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	preRect, _ = ebiten.NewImage(screenWidth/2, screenHeight/2, ebiten.FilterDefault)
	preRect.Fill(color.NRGBA{0, 255, 0, 50})
	showPreRect = false

	updateImg()
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Mandelbrot")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
