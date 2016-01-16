package server

type Server interface {
	Run() bool;
	Stop() bool;
}