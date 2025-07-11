package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mulaydm10/backend/internal/tools/agecalc"
	"github.com/mulaydm10/backend/internal/tools/audioconverter"
	"github.com/mulaydm10/backend/internal/tools/audiotrimmer"
	"github.com/mulaydm10/backend/internal/tools/base64coder"
	"github.com/mulaydm10/backend/internal/tools/bmi"
	"github.com/mulaydm10/backend/internal/tools/bmicalculator"
	"github.com/mulaydm10/backend/internal/tools/blogsection"
	"github.com/mulaydm10/backend/internal/tools/colorpicker"
	"github.com/mulaydm10/backend/internal/tools/emicalculator"
	"github.com/mulaydm10/backend/internal/tools/imagecompressor"
	"github.com/mulaydm10/backend/internal/tools/imageconverter"
	"github.com/mulaydm10/backend/internal/tools/imagecropper"
	"github.com/mulaydm10/backend/internal/tools/jsonformatter"
	"github.com/mulaydm10/backend/internal/tools/passwordgenerator"
	"github.com/mulaydm10/backend/internal/tools/qrcodegenerator"
	"github.com/mulaydm10/backend/internal/tools/sipcalculator"
	"github.com/mulaydm10/backend/internal/tools/speechtotext"
	"github.com/mulaydm10/backend/internal/tools/texttospeech"
	"github.com/mulaydm10/backend/internal/tools/timerstopwatch"
	"github.com/mulaydm10/backend/internal/tools/unitconverter"
	"github.com/mulaydm10/backend/internal/tools/videoconverter"
	"github.com/mulaydm10/backend/internal/tools/wordcounter"
)

func RegisterToolRoutes(router fiber.Router) {
	tools := router.Group("/tools")
	tools.Get("/bmi", bmi.Handler)
	tools.Get("/age", agecalc.Handler)
	tools.Get("/audioconverter", audioconverter.Handler)
	tools.Get("/audiotrimmer", audiotrimmer.Handler)
	tools.Get("/base64coder", base64coder.Handler)
	tools.Get("/bmicalculator", bmicalculator.Handler)
	tools.Get("/blogsection", blogsection.Handler)
	tools.Get("/colorpicker", colorpicker.Handler)
	tools.Get("/emicalculator", emicalculator.Handler)
	tools.Get("/imagecompressor", imagecompressor.Handler)
	tools.Get("/imageconverter", imageconverter.Handler)
	tools.Get("/imagecropper", imagecropper.Handler)
	tools.Get("/jsonformatter", jsonformatter.Handler)
	tools.Get("/passwordgenerator", passwordgenerator.Handler)
	tools.Get("/qrcodegenerator", qrcodegenerator.Handler)
	tools.Get("/sipcalculator", sipcalculator.Handler)
	tools.Get("/speechtotext", speechtotext.Handler)
	tools.Get("/texttospeech", texttospeech.Handler)
	tools.Get("/timerstopwatch", timerstopwatch.Handler)
	tools.Get("/unitconverter", unitconverter.Handler)
	tools.Get("/videoconverter", videoconverter.Handler)
	tools.Get("/wordcounter", wordcounter.Handler)
}
