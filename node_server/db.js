
const { v4: uuidv4 } = require("uuid");

const data=[
    {
        id:uuidv4(),
        title:"first notification",
        content:"Initial notification",
        time_stamp:new Date().getTime(),
    }
]


module.exports=data