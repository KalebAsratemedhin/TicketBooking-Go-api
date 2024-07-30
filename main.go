package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Ticket struct{
	ID int `json:"id"`
	Reason string `json:"reason"`
	Event string `json:"event"`
	Price float32 `json:"price"`
}

var tickets = []Ticket{
	{ID: 1, Reason: "Hello", Event: "Conference", Price: 34},
	{ID: 2, Reason: "Mello", Event: "Wedding", Price: 14},

}

func getTickets(c *gin.Context){
	
	c.JSON(http.StatusOK, tickets)
}

func createTicket(c *gin.Context){
	var newTicket Ticket

	if err:= c.ShouldBindJSON(&newTicket); err != nil{
		c.JSON(http.StatusBadRequest, newTicket)
		return
	}

	tickets = append(tickets, newTicket)
	c.JSON(http.StatusCreated, newTicket)
}

func getTicketById(c *gin.Context){
	ticketId, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	for _,ticket := range(tickets){
		if ticket.ID == int(ticketId){
			c.JSON(http.StatusFound, ticket)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error":"No such ticket"})

}

func main(){
	router := gin.Default()
	router.GET("/tickets", getTickets)
	router.POST("/tickets", createTicket)
	router.GET("/tickets/:id", getTicketById)

	router.Run()
}