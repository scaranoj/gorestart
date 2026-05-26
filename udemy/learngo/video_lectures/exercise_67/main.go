package main

//create a mockdatastore and a service so new entries can be added to it and lookups can also be performed

// User represents a user record in a database
type User struct {
	ID    int
	First string
}

// MockDatastore is for holding User records
type MockDatastore struct {
	Users map[int]User
}

// creata a method attached to MockDatastore that saves a user
// I cant't seem to quickly recall the input parms
// When wanting to save a user, you want to supply the name and write to a map. The id will be whatever their index in the map is
// before writing to the map, you need to make sure that someone is not already there
// use the comma okay idiom. If okay is true, then error, otherwise return error = nil and write to the map
// I intially tried to use (u User.First) as the input parm when all I needed was User. User.First is not a type, it's a field. Input parms need a type.
// I tried to just use `u` for the map lookup to check if user exists but actually needed to use the field. `u` is just a local var for the User type. You want to use the field ID
// to see if it the user is already there
// You forgot to add an if after the comma ok. If ok, the that means the user is already there and should throw an error
// So is was weird to me but the reason he's using the ID to check is because it's assumed that when a new user is created, both fields in the struct (First AND ID) will be provided
// /the comma ok will simply to a lookup using the usual map lookup syntax `[u.ID]`
func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		u.Errorf("User  %s already exists", u)
	}
	//md.Users[u.ID] = u

	//syntax example for writing to a map: `name[key] = val` e.g. `m["k1"] = 7`
	//Why did he only specify the ID, don't we need to set both an ID AND First? A: NO, becuase the value that is passed to this method bundles both (i.e. "User" is a struct type that
	//contains both fields/types/values. The end user will create a struct literal with field values that get passed to the SaveUser method). The only reason we are using ID here is
	//because we are writing to a map and need to specify the key? If I wanted to write a key value pair to the Users database, then I would use the syntax: Users[int]User, or since
	//we're in a func, use the mock database type receiver name `md` and input parm name `u`: md.Users[int]User, which is md.Users[u.ID] = u. I'm still confused. 'u' is the input parm
	//that was provided. A: yeah....so.....it's a struct type literal with 2 field values. Q: Okay, then wtf are we using u.ID as a key? Something feels redundant here (i.e. two keys of ID)
	md.Users[u.ID] = u
	return nil

}

// creata a method attached to MockDatastore that performs a user lookup
func (md MockDatastore) GetUser(u u.first) (User, error) {

	if err != ok {

	}
	return
}
