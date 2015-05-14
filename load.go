package main

var (
	// The only instance of loadState we need
	mainLoadState = new(loadState)
)

// The state for loading the game's assets. Implements fsm.State and mainState.
type loadState struct {
	totalAssets  int
	assetsLoaded int
	loadChannel  chan string
}

func (s *loadState) Name() string {
	return "loadState"
}

func (s *loadState) OnEnter() {
	println("loadState.OnEnter")
	initTest()

	if s.loadChannel == nil {
		s.loadChannel = make(chan string)
	}

	s.totalAssets = len(shaderAssets)
	s.totalAssets += len(textureAssets)
	// s.totalAssets += len(soundAssets)
	// s.totalAssets += len(modelAssets)

	for _, asset := range shaderAssets {
		go loadShaderAsset(&asset, s.loadChannel)
	}

	for _, asset := range textureAssets {
		go loadTextureAsset(&asset, s.loadChannel)
	}
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	// TODO: closing this causes an exception to be thrown about sending to closed
	//       channel by who ever is the last to load.
	// close(s.loadChannel)
}

func (s *loadState) Update() {
	select {
	case loaded := <-s.loadChannel:
		println("loaded", loaded)
		s.assetsLoaded += 1
	default:
	}

	if s.assetsLoaded == s.totalAssets {
		mainSm.GotoState(mainMenuState)
		return
	}
}

func (s *loadState) Draw() {
	// percent := float64(s.assetsLoaded) / float64(s.totalAssets) * 100.0
	// println("loading... ", percent, "%")
	drawTest()
}
