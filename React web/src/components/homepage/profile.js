import React, { useState, useEffect } from "react";
import axios from "axios";
import './homepage.css'
const Profile = () => {
    const [profileName, setProfileName] = useState("");
    const [profileRole, setProfileRole] = useState("");
    const [profileEmail, setProfileEmail] = useState("");


    const profileData = async () => {
        try {
            const res = await axios.get("http://localhost:8000/user", { withCredentials: true });
            console.log(res.data);
            setProfileEmail(res.data.email);
            setProfileRole(res.data.role);
            setProfileName(res.data.username);
        } catch (error) {
            console.log(error);
        }
    };

    useEffect(() => {
        profileData();
    });

    return (
        <div onLoad={() => profileData()}>
            <div className="card"><div className="card" style={{ width: "18rem" }}>
                <img src={require('../img/images.png')} className="card-img-top" alt="employee" style={{ width: "100%" }} ></img>
                <div className="card-body">
                    <h4 className="card-title">Name: {profileName}</h4>
                    <p className="card-text">Email: {profileEmail}</p>
                    <p className="card-text">Role: {profileRole}</p>
                </div>
            </div>
            </div>
        </div>
    );
};

export default Profile;