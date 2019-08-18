import React, { Component } from 'react'
import Axios from 'axios'
import { Link } from 'react-router-dom';
export class Student extends Component {
    constructor(props) {
        super(props)

        this.state = { rows: [] }
    }
    getStudent = () => {
        var dataAllStudent = []
        Axios.get('http://localhost:8080/students').then(result => {
            console.log(result.data)
            result.data.forEach(item => {
                dataAllStudent.push(item)
            })
            this.setState({ rows: dataAllStudent })
        })
    }
    componentDidMount() {
        { this.getStudent() }
    }

    render() {
        return (

            <div>
                <center>
                    <h1>Students</h1>
                    <table border='1' width='80%'>
                        <tbody>
                            <tr><th>No.</th>
                                <th>Student ID</th>
                                <th>First Name</th>
                                <th>Last Name</th>
                                <th>Major</th>
                            </tr>
                            {this.state.rows.map((item, index) => (
                                <tr key={index}><td>{index+1}</td>
                                    <td><Link to={`/student/${item.student_id}`}>{item.student_code}</Link></td>
                                    <td>{item.first_name}</td>
                                    <td>{item.last_name}</td>
                                    <td>{item.Major.major_name}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </center>
            </div>
        )
    }

}

export default Student
