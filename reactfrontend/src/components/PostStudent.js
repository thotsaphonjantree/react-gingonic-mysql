import React, { Component } from 'react'
import axios from 'axios';

export default class PostStudent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            student_code: '', first_name: '', last_name: '', major_id: '', rows: ['']
        }
        this.handleInputChange = this.handleInputChange.bind(this);
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
        this.setState({
            [name]: value,
        });
    }





    componentDidMount() {

        { this.getMajor() }
    }
    getMajor = () => {
        var dataAllMajor = []
        axios.get('http://localhost:8080/majors').then(result => {
            // console.log(result.data)
            result.data.forEach(item => {
                dataAllMajor.push(item)
            })
            this.setState({ rows: dataAllMajor })
        })
    }

    // handleChange = event => {
    //     const name = event.target.name;
    //     const value = event.target.value;
    //     this.setState({
    //         ...this.state, [name]: { ...this.state[name], value }
    //     });
    // }

    handleSubmit = event => {
        event.preventDefault();
        axios.post(`http://localhost:8080/student`, {
            student_code: this.state.student_code,
            first_name: this.state.first_name,
            last_name: this.state.last_name,
            major_id: parseInt(this.state.major_id)
        })
            .then(res => {
                console.log(res);
                console.log(res.data);
            })
        console.dir(this.state);
        window.location = '/students/';
    }
    render() {
        return (
            <div><center>
                <h1>Add Student</h1>
                <form onSubmit={this.handleSubmit}>
                    <label>
                        Student ID:
                <input type="text" name="student_code" value={this.state.student_code} onChange={this.handleInputChange} />
                    </label><br />
                    <label>
                        First Name:
                <input type="text" name="first_name" value={this.state.first_name} onChange={this.handleInputChange} />
                    </label><br />
                    <label>
                        Last Name:
                <input type="text" name="last_name" value={this.state.last_name} onChange={this.handleInputChange} />
                    </label><br />
                    <label>
                        <select name="major_id" value={this.state.major_id} onChange={this.handleInputChange}>
                            <option value="DEFAULT">Select Major:</option>
                            {this.state.rows.map((item, index) => (
                                <option value={item.major_id} key={index} >{item.major_name}</option>
                            ))}
                        </select>
                    </label>
                    <br />
                    <button type="submit">Add</button>
                </form>
            </center>


            </div>
        )
    }
}
