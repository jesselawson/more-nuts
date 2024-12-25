components {
  id: "tree_spawner"
  component: "/actors/trees/tree_spawner/tree_spawner.script"
}
embedded_components {
  id: "treefactory"
  type: "factory"
  data: "prototype: \"/actors/trees/tree_spawner/tree_proto.go\"\n"
  ""
}
embedded_components {
  id: "nutfactory"
  type: "factory"
  data: "prototype: \"/actors/nut/nut_proto.go\"\n"
  ""
}
