package api

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "UserProfile",
        "GET",
        "/userprofile",
        GetUserProfile,
    },
    Route{
        "UserProfile",
        "PUT",
        "/userprofile",
        EditUserProfile
    },
    Route{
        "Login"
        "GET"
        "/login"
        Login,
    },
    Route{
        "Logout"
        "GET"
        "/logout"
        Logout,
    },
}
