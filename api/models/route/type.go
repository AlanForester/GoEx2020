package route

type RouteType int

const (
	ROUTE_TYPE_GET RouteType = iota
	ROUTE_TYPE_POST
	ROUTE_TYPE_PUT
	ROUTE_TYPE_PATCH
	ROUTE_TYPE_DELETE
)
