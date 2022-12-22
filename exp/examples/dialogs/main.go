package main

import (
	_ "embed"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/wailsapp/wails/exp/pkg/application"
)

func main() {

	app := application.New()
	app.SetName("Dialogs Demo")
	app.SetDescription("A demo of the Wails dialogs")

	// Create a custom menu
	menu := app.NewMenu()
	menu.AddRole(application.AppMenu)

	// Let's make a "Demo" menu
	infoMenu := menu.AddSubmenu("Info")
	infoMenu.Add("Info").OnClick(func(ctx *application.Context) {
		dialog := app.NewInfoDialog()
		dialog.SetTitle("Custom Title")
		dialog.SetMessage("This is a custom message")
		dialog.Show()
	})
	infoMenu.Add("Info (Title only)").OnClick(func(ctx *application.Context) {
		dialog := app.NewInfoDialog()
		dialog.SetTitle("Custom Title")
		dialog.Show()
	})
	infoMenu.Add("Info (Message only)").OnClick(func(ctx *application.Context) {
		dialog := app.NewInfoDialog()
		dialog.SetMessage("This is a custom message")
		dialog.Show()
	})
	infoMenu.Add("Info (Custom Icon)").OnClick(func(ctx *application.Context) {
		dialog := app.NewInfoDialog()
		dialog.SetTitle("Custom Icon Example")
		dialog.SetMessage("Using a custom icon")
		dialog.SetIcon(application.DefaultApplicationIcon)
		dialog.Show()
	})

	questionMenu := menu.AddSubmenu("Question")
	questionMenu.Add("Question (No default)").OnClick(func(ctx *application.Context) {
		dialog := app.NewQuestionDialog()
		dialog.SetMessage("No default button")
		dialog.AddButton("Yes")
		dialog.AddButton("No")
		dialog.Show()
	})
	questionMenu.Add("Question (With Default)").OnClick(func(ctx *application.Context) {
		dialog := app.NewQuestionDialog()
		dialog.SetTitle("Quit")
		dialog.SetMessage("You have unsaved work. Are you sure you want to quit?")
		dialog.AddButton("Yes").OnClick(func() {
			app.Quit()
		})
		no := dialog.AddButton("No")
		dialog.SetDefaultButton(no)
		dialog.Show()
	})
	questionMenu.Add("Question (With Cancel)").OnClick(func(ctx *application.Context) {
		dialog := app.NewQuestionDialog().
			SetTitle("Update").
			SetMessage("The cancel button is selected when pressing escape")
		download := dialog.AddButton("📥 Download")
		download.OnClick(func() {
			app.NewInfoDialog().SetMessage("Downloading...").Show()
		})
		no := dialog.AddButton("Cancel")
		dialog.SetDefaultButton(download)
		dialog.SetCancelButton(no)
		dialog.Show()
	})
	questionMenu.Add("Question (Custom Icon)").OnClick(func(ctx *application.Context) {
		dialog := app.NewQuestionDialog()
		dialog.SetTitle("Custom Icon Example")
		dialog.SetMessage("Using a custom icon")
		dialog.SetIcon(application.WailsLogoWhiteTransparent)
		dialog.SetDefaultButton(dialog.AddButton("I like it!"))
		dialog.AddButton("Not so keen...")
		dialog.Show()
	})

	warningMenu := menu.AddSubmenu("Warning")
	warningMenu.Add("Warning").OnClick(func(ctx *application.Context) {
		dialog := app.NewWarningDialog()
		dialog.SetTitle("Custom Title")
		dialog.SetMessage("This is a custom message")
		dialog.Show()
	})
	warningMenu.Add("Warning (Title only)").OnClick(func(ctx *application.Context) {
		dialog := app.NewWarningDialog()
		dialog.SetTitle("Custom Title")
		dialog.Show()
	})
	warningMenu.Add("Warning (Message only)").OnClick(func(ctx *application.Context) {
		dialog := app.NewWarningDialog()
		dialog.SetMessage("This is a custom message")
		dialog.Show()
	})
	warningMenu.Add("Warning (Custom Icon)").OnClick(func(ctx *application.Context) {
		dialog := app.NewWarningDialog()
		dialog.SetTitle("Custom Icon Example")
		dialog.SetMessage("Using a custom icon")
		dialog.SetIcon(application.DefaultApplicationIcon)
		dialog.Show()
	})

	errorMenu := menu.AddSubmenu("Error")
	errorMenu.Add("Error").OnClick(func(ctx *application.Context) {
		dialog := app.NewErrorDialog()
		dialog.SetTitle("Ooops")
		dialog.SetMessage("I accidentally the whole of Twitter")
		dialog.Show()
	})
	errorMenu.Add("Error (Title Only)").OnClick(func(ctx *application.Context) {
		dialog := app.NewErrorDialog()
		dialog.SetTitle("Custom Title")
		dialog.Show()
	})
	errorMenu.Add("Error (Custom Message)").OnClick(func(ctx *application.Context) {
		dialog := app.NewErrorDialog()
		dialog.SetMessage("This is a custom message")
		dialog.Show()
	})
	errorMenu.Add("Error (Custom Icon)").OnClick(func(ctx *application.Context) {
		dialog := app.NewErrorDialog()
		dialog.SetTitle("Custom Icon Example")
		dialog.SetMessage("Using a custom icon")
		dialog.SetIcon(application.WailsLogoWhite)
		dialog.Show()
	})

	openMenu := menu.AddSubmenu("Open")
	openMenu.Add("Open File").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseFiles(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file selected").Show()
		}
	})
	openMenu.Add("Open File (Show Hidden Files)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseFiles(true).
			CanCreateDirectories(true).
			ShowHiddenFiles(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file selected").Show()
		}
	})
	openMenu.Add("Open File (Attach to window)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseFiles(true).
			CanCreateDirectories(true).
			ShowHiddenFiles(true).
			AttachToWindow(app.GetCurrentWindow()).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file selected").Show()
		}
	})
	openMenu.Add("Open Multiple Files (Show Hidden Files)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseFiles(true).
			CanCreateDirectories(true).
			ShowHiddenFiles(true).
			PromptForMultipleSelection()
		if len(result) > 0 {
			app.NewInfoDialog().SetMessage(strings.Join(result, ",")).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file selected").Show()
		}
	})
	openMenu.Add("Open Directory").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseDirectories(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No directory selected").Show()
		}
	})
	openMenu.Add("Open Directory (Create Directories)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseDirectories(true).
			CanCreateDirectories(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No directory selected").Show()
		}
	})
	openMenu.Add("Open Directory (Resolves Aliases)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewOpenFileDialog().
			CanChooseDirectories(true).
			CanCreateDirectories(true).
			ResolvesAliases(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No directory selected").Show()
		}
	})
	openMenu.Add("Open File/Directory (Set Title)").OnClick(func(ctx *application.Context) {
		dialog := app.NewOpenFileDialog().
			CanChooseDirectories(true).
			CanCreateDirectories(true).
			ResolvesAliases(true)
		if runtime.GOOS == "darwin" {
			dialog.SetMessage("Select a file/directory")
		} else {
			dialog.SetTitle("Select a file/directory")
		}

		result, _ := dialog.PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file/directory selected").Show()
		}
	})
	openMenu.Add("Open (Full Example)").OnClick(func(ctx *application.Context) {
		cwd, _ := os.Getwd()
		dialog := app.NewOpenFileDialog().
			SetTitle("Select a file").
			SetMessage("Select a file to open").
			SetButtonText("Let's do this!").
			SetDirectory(cwd).
			CanCreateDirectories(true).
			ResolvesAliases(true).
			AllowsOtherFileTypes(true).
			TreatsFilePackagesAsDirectories(true).
			ShowHiddenFiles(true).
			CanSelectHiddenExtension(true).
			AddFilter("Text Files", "*.txt; *.md").
			AddFilter("Video Files", "*.mov; *.mp4; *.avi")

		if runtime.GOOS == "darwin" {
			dialog.SetMessage("Select a file")
		} else {
			dialog.SetTitle("Select a file")
		}

		result, _ := dialog.PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		} else {
			app.NewInfoDialog().SetMessage("No file selected").Show()
		}
	})

	saveMenu := menu.AddSubmenu("Save")
	saveMenu.Add("Select File (Defaults)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewSaveFileDialog().
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		}
	})
	saveMenu.Add("Select File (Attach To Window)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewSaveFileDialog().
			AttachToWindow(app.GetCurrentWindow()).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		}
	})
	saveMenu.Add("Select File (Show Hidden Files)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewSaveFileDialog().
			ShowHiddenFiles(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		}
	})
	saveMenu.Add("Select File (Cannot Create Directories)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewSaveFileDialog().
			CanCreateDirectories(false).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		}
	})
	saveMenu.Add("Select File (Full Example)").OnClick(func(ctx *application.Context) {
		result, _ := app.NewSaveFileDialog().
			CanCreateDirectories(false).
			ShowHiddenFiles(true).
			SetMessage("Select a file").
			SetDirectory("/Applications").
			SetButtonText("Let's do this!").
			SetFilename("README.md").
			HideExtension(true).
			AllowsOtherFileTypes(true).
			TreatsFilePackagesAsDirectories(true).
			ShowHiddenFiles(true).
			PromptForSingleSelection()
		if result != "" {
			app.NewInfoDialog().SetMessage(result).Show()
		}
	})

	app.SetMenu(menu)

	app.NewWindow()
	app.NewWindow()

	err := app.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}