embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/actors/trees/tree_spawner/tree02.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/builtins/materials/model.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/actors/trees/tree_spawner/Texture-colors.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    x: 1.0
    y: 1.0
  }
  rotation {
    x: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "collider"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -0.5\n"
  "      z: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 0.25\n"
  "  data: 0.25\n"
  "  data: 0.5\n"
  "}\n"
  ""
}
