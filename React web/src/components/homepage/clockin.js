import React, { useState } from 'react'
import './homepage.css'
import { Switch } from 'antd';

import axios from "axios"


export default function Clockin({ state, handleSwitchChange }) {

    const [toggle, setToggle] = useState(false);

    const toggler = () => {
        toggle ? setToggle(false) : setToggle(true);
        axios.get("http://localhost:8000/clockin")
            .then(res => {
                alert(res.data.message)
            })
    }
    return (
        <div >
            <Switch className='clockButton' checkedChildren="Clocked In" unCheckedChildren="Clocked Out" onClick={toggler} />
        </div>
    )
}

