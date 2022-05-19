import React from 'react'
import "./homepage/homepage.css"

export default function Header({ setLoginUser }) {
    return (
        <div>
            <nav className="navbar bg-primary" style={{ position: "static" }}>
                <div className="container-fluid position-relative">
                    <div className="navbar-brand">
                        <img src={require('./img/mckinley_logo.png')} alt="logo" width="30" height="24" className="d-inline-block align-text-top " />
                        <h2 className="text-light text-center d-inline px-5 ">MnR - EasyHR </h2>
                    </div>
                    <button className="btn btn-outline-danger" onClick={() => setLoginUser({})} type="submit">Log Out</button>

                </div>
            </nav>
        </div>
    )
}
