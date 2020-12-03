package main

type Cfg struct {
	Original   *Node      `json:"original"`
	ProPkgPath string     `json:"pro_pkg_path"`
	ProNodes   []*ProNode `json:"pro_nodes"`
}

type Node struct {
	Addr      string `json:"addr"`
	SSHUser   string `json:"ssh_user"`
	SSHPasswd string `json:"ssh_passwd"`
}

type ProNode struct {
	Hostname  string `json:"hostname"`
	Ip        string `json:"ip"`
	SSHUser   string `json:"ssh_user"`
	SSHPasswd string `json:"ssh_passwd"`
}