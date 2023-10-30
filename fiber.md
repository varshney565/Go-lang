Returning a response

c *fiber.Ctx

c.Status(fiber.StatusOk)                  		200 direct
c.status(fiber.StatusNotFound)       		404
c.status(fiber.StatusForbidden)       		403
c.status(fiber.StatusUnauthorized)  		401 
c.status(fiber.StatusInternalServerError)	500
c.status(fiber.StatusBadRequest)   		400

c.SendString(response)

type Category struct {
    Name  string `json:"name"`
    Count int    `json:"count"`
}
category := Category{
     Name:  "Example Category",
     Count: 10,
}

c.JSON(category)

or 

c.status(200).JSON(fiber.map{“status” : “success”,”message” : “Not found”,”data” : nil})



Getting from the header


header-val := c.Get(“header-passed-value”)

intValue, err := strconv.Atoi(headerValue)
if err != nil {
   return c.Status(fiber.StatusBadRequest).SendString("Invalid header value")
}


Getting from path param and query param


name := c.Query(“name”)
id := c.Param(“id”)




Getting data from the post-call

payload := struct {
     Name  string `json:"name"`
     Email string `json:"email"`
}{}



if err := c.BodyParser(&data); err != nil {
    return err
}



Config


func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}



Creating a route 

func SetupRoutes(app *fiber.App) {
	user := api.Group("/user")
	user.Post("/create",	middleware,	logic)
	user.Get(“/:id”, logic)
	user.Post()
	user.Put()
	user.Delete()
}

Middleware

func verify(c *fiber.Ctx) error {
	//anything that is needed to verify
	
	return c.Next()
}


Controller

func GetSingleUser(c *fiber.Ctx) error {
       //Buisness Logic
}


