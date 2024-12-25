local Constants = {}

function Constants.alter_state(the_state, v)
	the_state.value = the_state.value + v
end

function Constants.get_state(the_state)
	return the_state.value
end

function Constants.new(v)
	local state = {
		value = v
	}
	return state
end

return Constants