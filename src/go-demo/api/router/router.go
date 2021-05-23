package router

import (
    PlayerController "go-demo/api/controller"
    LeagueController "go-demo/api/controller"
    TournamentController "go-demo/api/controller"
    MatchupController "go-demo/api/controller"
    GameController "go-demo/api/controller"
    jsonResponseFormatter "go-demo/api/response"
    "github.com/gorilla/mux"
    
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()    


	//----------------------------------------------------------------------------
    // Player Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard CRUD Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/Player/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.GetPlayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Player", jsonResponseFormatter.FormatToJSON(PlayerController.GetAllPlayer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPlayer", jsonResponseFormatter.FormatToJSON(PlayerController.CreatePlayer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Player/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.UpdatePlayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePlayer/{id}", jsonResponseFormatter.FormatToJSON(PlayerController.DeletePlayer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------    

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------    

	//----------------------------------------------------------------------------
    // League Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard CRUD Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/League/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.GetLeague)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/League", jsonResponseFormatter.FormatToJSON(LeagueController.GetAllLeague)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLeague", jsonResponseFormatter.FormatToJSON(LeagueController.CreateLeague)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/League/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.UpdateLeague)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLeague/{id}", jsonResponseFormatter.FormatToJSON(LeagueController.DeleteLeague)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------    

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/AddPlayers/{parentId}/playersId", jsonResponseFormatter.FormatToJSON(LeagueController.AddPlayers)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlayers/{parentId}/playersIds", jsonResponseFormatter.FormatToJSON(LeagueController.RemovePlayers)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Tournament Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard CRUD Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/Tournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.GetTournament)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Tournament", jsonResponseFormatter.FormatToJSON(TournamentController.GetAllTournament)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTournament", jsonResponseFormatter.FormatToJSON(TournamentController.CreateTournament)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Tournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.UpdateTournament)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTournament/{id}", jsonResponseFormatter.FormatToJSON(TournamentController.DeleteTournament)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------    

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/AddMatchups/{parentId}/matchupsId", jsonResponseFormatter.FormatToJSON(TournamentController.AddMatchups)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMatchups/{parentId}/matchupsIds", jsonResponseFormatter.FormatToJSON(TournamentController.RemoveMatchups)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Matchup Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard CRUD Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/Matchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.GetMatchup)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Matchup", jsonResponseFormatter.FormatToJSON(MatchupController.GetAllMatchup)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMatchup", jsonResponseFormatter.FormatToJSON(MatchupController.CreateMatchup)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Matchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.UpdateMatchup)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMatchup/{id}", jsonResponseFormatter.FormatToJSON(MatchupController.DeleteMatchup)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------    

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/AddGames/{parentId}/gamesId", jsonResponseFormatter.FormatToJSON(MatchupController.AddGames)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGames/{parentId}/gamesIds", jsonResponseFormatter.FormatToJSON(MatchupController.RemoveGames)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
    // Game Routes to JSON response formatter first 
    // then to the correct Controller function
	//----------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	// Standard CRUD Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/Game/{id}", jsonResponseFormatter.FormatToJSON(GameController.GetGame)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Game", jsonResponseFormatter.FormatToJSON(GameController.GetAllGame)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGame", jsonResponseFormatter.FormatToJSON(GameController.CreateGame)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Game/{id}", jsonResponseFormatter.FormatToJSON(GameController.UpdateGame)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGame/{id}", jsonResponseFormatter.FormatToJSON(GameController.DeleteGame)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Single Association Routers
	//----------------------------------------------------------------------------    
    router.HandleFunc("/api/AssignMatchup/{parentId}/matchupId", jsonResponseFormatter.FormatToJSON(GameController.AssignMatchup)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMatchup/{parentId}", jsonResponseFormatter.FormatToJSON(GameController.UnassignMatchup)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlayer/{parentId}/playerId", jsonResponseFormatter.FormatToJSON(GameController.AssignPlayer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlayer/{parentId}", jsonResponseFormatter.FormatToJSON(GameController.UnassignPlayer)).Methods("DELETE", "OPTIONS")

	//----------------------------------------------------------------------------
	// Multiple Association Routers
	//----------------------------------------------------------------------------    
    return router
}