components {
  id: "home_tree"
  component: "/actors/trees/home_tree/home_tree.script"
}
embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/actors/trees/home_tree/tree_07.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/actors/trees/home_tree/home_tree.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/actors/trees/tree_spawner/Texture-colors.png\"\n"
  "  }\n"
  "}\n"
  ""
  rotation {
    x: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "canopy"
  type: "model"
  data: "mesh: \"/actors/trees/tree_spawner/tree_01.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/builtins/materials/model_instanced.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/actors/trees/tree_spawner/canopy_home.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: -1.0
    y: -2.0
    z: 13.0
  }
  rotation {
    x: 0.07376538
    y: 0.69605637
    z: 0.36769623
    w: 0.6122612
  }
}
embedded_components {
  id: "canopy1"
  type: "model"
  data: "mesh: \"/actors/trees/tree_spawner/tree_01.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/builtins/materials/model_instanced.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/actors/trees/tree_spawner/canopy_home.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: -1.0
    y: -4.0
    z: 11.0
  }
  rotation {
    x: 0.24779275
    y: 0.0747399
    z: 0.2789331
    w: 0.9247751
  }
}
embedded_components {
  id: "canopy2"
  type: "model"
  data: "mesh: \"/actors/trees/tree_spawner/tree_01.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/builtins/materials/model_instanced.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/actors/trees/tree_spawner/canopy_home.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: -1.0
    y: -4.0
    z: 17.0
  }
  rotation {
    x: -0.25096285
    y: -0.07569608
    z: 0.27867514
    w: 0.92391986
  }
}
embedded_components {
  id: "collider"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"hometree\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -2.5\n"
  "      y: -0.78\n"
  "      z: 7.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: -0.31730467\n"
  "      w: 0.94832367\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 2.5\n"
  "  data: 3.25\n"
  "  data: 2.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "plusonenutfactory"
  type: "factory"
  data: "prototype: \"/actors/trees/home_tree/plus_one_nut_proto.go\"\n"
  ""
}
