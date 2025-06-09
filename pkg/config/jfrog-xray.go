package config

func LoadConfig() {

}

func JfrogLogin() {

}

// var artifactoryURL = "https://your.remote.server/artifactory/"

// rtDetails := auth.NewArtifactoryDetails()
// rtDetails.SetUrl(artifactoryURL)
// fmt.Printf("Artifactory login name....\n")
// var userName string
// userName = [some way of getting user name eg via command line]
// fmt.Printf("Artifactory password....\n")
// var passWord string
// passWord = [some way of getting artifactory password eg via command line]
// rtDetails.SetUser(userName)
// rtDetails.SetPassword(passWord)

// artifactoryHTTP := &http.Client{Transport: &http.Transport{}}
// serviceConfig, err := config.NewConfigBuilder().SetServiceDetails(rtDetails).SetDryRun(false).SetHttpClient(artifactoryHTTP).Build()
// if err != nil {
//     panic(err)
// }
// rtManager, errSC := artifactory.New(serviceConfig)
// if errSC != nil {
//     panic(errSC)
// }
// params := services.NewDownloadParams()
// params.Pattern = "[repo_name]/path/to/your/file/your_file"
// _, _, err = rtManager.DownloadFiles(params)
// if err != nil {
//     fmt.Printf("%s\n", err.Error())
//     panic(err)
// }
