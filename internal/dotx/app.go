package dotx

import (
	"fmt"
	"github.com/mlang97/dotx/internal/config"
	"github.com/mlang97/dotx/internal/fs"
	"path/filepath"
)

type App struct {
	AppConfig  config.AppConfig
	repoConfig config.RepoConfig
}

func New(appConfig config.AppConfig, repoConfig config.RepoConfig) App {
	return App{
		AppConfig:  appConfig,
		repoConfig: repoConfig,
	}
}

func (a App) AddDotfile(path string) error {
	filename := filepath.Base(path)
	source := fs.NewPath(path)
	dest := fs.NewPath(filepath.Join(config.RepoDirPath(), filename))

	fmt.Println(path)
	fmt.Println(filename)
	fmt.Println(source)
	fmt.Println(dest)

	//if a.repoConfig.GetDotfile(source.AbsPath()) != (config.Dotfile{}) {
	//	return errors.New("dotfile already exist")
	//}
	//
	//if err := fs.Move(source, dest); err != nil {
	//	return err
	//}
	//
	//if err := fs.Symlink(dest, source); err != nil {
	//	return err
	//}
	//
	//// normalize paths
	//home, _ := os.UserHomeDir()
	//sourcePath := strings.Replace(source.AbsPath(), home, "$HOME", 1)
	//destinationPath := strings.Replace(dest.AbsPath(), a.appConfig.RepoDir, "", 1)
	//
	//dotfile := config.Dotfile{
	//	Source:      destinationPath,
	//	Destination: sourcePath,
	//}
	//
	//if err := a.repoConfig.WriteDotfile(dotfile); err != nil {
	//	return errors.New("failed to write config")
	//}

	return nil
}

func (a App) DeployDotfiles() error {
	// for _, dotfile := range a.repoConfig.Dotfiles {
	// 	sourcePath := filepath.Join(a.appConfig.RepoDir, dotfile.Source)
	// 	destPath := os.ExpandEnv(dotfile.Destination)

	// 	_, err := os.Stat(destPath)
	// 	if err == nil {
	// 		symlinkPath, err := a.fs.SymlinkPath(destPath)
	// 		if err == nil && sourcePath == symlinkPath {
	// 			a.Logger.Warn("file already deployed from dotx")
	// 			continue
	// 		}

	// 		reader := bufio.NewReader(os.Stdin)
	// 		fmt.Print("Dotfile already exists on your sytem. Do you want to backup and overwrite? (Y/n): ")
	// 		char, _, err := reader.ReadRune()
	// 		if err != nil {
	// 			return err
	// 		}

	// 		switch char {
	// 		case 'y', 'Y', '\n':
	// 		// TODO move existing file to backup folder
	// 		case 'n', 'N':
	// 			continue
	// 		default:
	// 			return errors.New("invalid user input")
	// 		}
	// 	}

	// 	dirPath := filepath.Dir(destPath)
	// 	if err := a.fs.Mkdir(dirPath); err != nil {
	// 		return errors.New("failed to create dir")
	// 	}

	// 	if err := a.fs.Symlink(sourcePath, destPath); err != nil {
	// 		return errors.New("failed to symlink file")
	// 	}
	// }

	return nil
}

func (a App) InitializeRemoteRepo(remoteRepo string) error {
	//command := exec.Command("git", "clone", remoteRepo, a.appConfig.RepoDir)
	//if err := command.Run(); err != nil {
	//	return errors.New("failed to clone remote repo")
	//}

	return nil
}
