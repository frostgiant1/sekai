package kira

import (
	"net/http"

	"github.com/KiraCore/sekai/INTERX/common"
	"github.com/KiraCore/sekai/INTERX/config"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterKiraTokensRoutes registers kira tokens query routers.
func RegisterKiraTokensRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryKiraTokensAliases, QueryKiraTokensAliasesRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryKiraTokensRates, QueryKiraTokensRatesRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryKiraTokensAliases, "This is an API to query all tokens aliases.", true)
	common.AddRPCMethod("GET", config.QueryKiraTokensRates, "This is an API to query all tokens rates.", true)
}

func queryKiraTokensAliasesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryKiraTokensAliasesRequest is a function to query all tokens aliases.
func QueryKiraTokensAliasesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-aliases] Entering token aliases query")

		if !common.RPCMethods["GET"][config.QueryKiraTokensAliases].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryKiraTokensAliases].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-aliases] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryKiraTokensAliasesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryKiraTokensAliases].CachingEnabled)
	}
}

func queryKiraTokensRatesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryKiraTokensRatesRequest is a function to query all tokens rates.
func QueryKiraTokensRatesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-rates] Entering token rates query")

		if !common.RPCMethods["GET"][config.QueryKiraTokensRates].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryKiraTokensRates].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-rates] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryKiraTokensRatesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryKiraTokensRates].CachingEnabled)
	}
}
