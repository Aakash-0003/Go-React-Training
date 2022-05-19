import React from "react"
import Header from "../Header"
import Clockin from "./clockin"
import "./homepage.css"
import Profile from "./profile"

const Homepage = () => {
    return (
        <div className="homepage">
            <div className="header">
                <Header />
            </div>
            <div className="clockComp" >
                <Clockin />
            </div>
            <div className="profile">
                <Profile />
            </div>


        </div>
    )
}


export default Homepage