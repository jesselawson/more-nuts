components {
  id: "bad_squirrel_spawner"
  component: "/actors/trees/home_tree/bad_squirrel_spawner.script"
}
embedded_components {
  id: "badsquirrelspawner"
  type: "factory"
  data: "prototype: \"/actors/trees/home_tree/bad_squirrel_proto.go\"\n"
  ""
}
embedded_components {
  id: "collider"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"badsquirrelspawner\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 50.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"mark_territory_overlay\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_SCREEN\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/actors/trees/home_tree/bad_squirrel_spawner.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
embedded_components {
  id: "1"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-01.ogg\"\n"
  ""
}
embedded_components {
  id: "2"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-02.ogg\"\n"
  ""
}
embedded_components {
  id: "3"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-03.ogg\"\n"
  ""
}
embedded_components {
  id: "4"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-04.ogg\"\n"
  ""
}
embedded_components {
  id: "5"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-05.ogg\"\n"
  ""
}
embedded_components {
  id: "6"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-06.ogg\"\n"
  ""
}
embedded_components {
  id: "7"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-07.ogg\"\n"
  ""
}
embedded_components {
  id: "8"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-08.ogg\"\n"
  ""
}
embedded_components {
  id: "9"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-09.ogg\"\n"
  ""
}
embedded_components {
  id: "10"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-10.ogg\"\n"
  ""
}
embedded_components {
  id: "11"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-11.ogg\"\n"
  ""
}
embedded_components {
  id: "12"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-12.ogg\"\n"
  ""
}
embedded_components {
  id: "13"
  type: "sound"
  data: "sound: \"/assets/sounds/squirrel-13.ogg\"\n"
  ""
}
