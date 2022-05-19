import React, { useState } from "react"
import "./login.css"
import axios from "axios"
import { useHistory } from "react-router-dom"

const axiosConfig = {
    headers: {
        'Content-Type': 'application/json;charset=UTF-8',
        "Access-Control-Allow-Origin": "*",
    }

};
const Login = ({ setLoginUser }) => {

    const history = useHistory()

    const [user, setUser] = useState({

        email: "",
        password: ""
    })

    const handleChange = e => {
        const { name, value } = e.target
        setUser({
            ...user,
            [name]: value
        })
    }

    const login = () => {
        axios.post("http://localhost:8000/login", user, axiosConfig)
            .then(res => {
                alert(res.data.message)
                setLoginUser(res.data.user)
                history.push("/homepage")

            })
    }

    return (
        <div>
            <div className="drop-down">
                <div className="heading text-center text h2">Welcome to MnR - EasyHR</div>
            </div>
            <div className="container-login">
                <div className="login">
                    <h1>Login</h1>
                    <input type="text" name="email" value={user.email} onChange={handleChange} placeholder="Enter your Email"></input>
                    <input type="password" name="password" value={user.password} onChange={handleChange} placeholder="Enter your Password" ></input>
                    <div className="button" onClick={login}>Login</div>
                    <div>or</div>
                    <div className="button" onClick={() => history.push("/register")}>Register</div>
                </div>
            </div>
        </div>
    )
}

export default Login