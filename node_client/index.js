const express = require("express");
const client = require("./client");

const app = express();

app.use(express.json({
    limit:"50mb",    
}))


app.get("/", (req, res) => {
    client.GetAllNotifications(null,(err,data)=>{
        if(!err) res.json(data.notification_list)
        else res.status(500).send(err)
    })  
})

app.get("/add/:title/:content", (req, res) => {
    client.AddNotifications(req.params,(err,data)=>{
        if(!err) res.json(data.notification_list)
        else res.status(500).send(err)
    })
})

app.get("/delete/:id", (req, res) => {
    client.RemoveNotifications(req.params,(err,data)=>{
        if(!err) res.send("ok")
        else res.status(500).send(err)
    })
})


const PORT = 3000;
app.listen(PORT, () => {
    console.log("Server running at port %d", PORT);
});