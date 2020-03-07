package main

import (
	dailyspend "MyDailySpendGoal/DailySpend"
	profile "MyDailySpendGoal/Profile"
	todo "MyDailySpendGoal/TODO"
	"net/http"
)

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
		todo.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		todo.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		todo.TodoShow,
	}, Route{
		"TodoCreate",
		"POST",
		"/todos",
		todo.TodoCreate,
	}, Route{
		"Index",
		"get",
		"/profile",
		profile.Index,
	},
	Route{
		"Index",
		"post",
		"/profile",
		profile.Create,
	},
	Route{
		"Get",
		"get",
		"/profile/Get",
		profile.Get,
	},
	Route{
		"Get",
		"get",
		"/DailySpend/Get",
		dailyspend.Get,
	},
}
