import React, { Component } from 'react'
import Axios from 'axios'

export class student extends Component {
    constructor(props) {
        super(props)

        this.state = { rows: [] }
    }
    myfunction = () => {
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
        { this.myfunction() }
    }

    render() {
        return (
            
            <div>
                <center>
                <table border='1' width='80%'>
                    <tbody>
                        <tr><th>Student ID</th>
                            <th>First Name</th>
                            <th>Last Name</th>
                            <th>Major</th>
                        </tr>
                        {this.state.rows.map(item => (
                            <tr key={item.student_id}><td>{item.student_code}</td>
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

export default student
