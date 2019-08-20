import React, { Component } from 'react'
import Axios from 'axios'
import { Link } from 'react-router-dom';

export class Major extends Component {
    constructor(props) {
        super(props)

        this.state = { rows: [] }
    }

    componentDidMount() {

        { this.getMajor() }
    }

    getMajor = () => {
        var dataAllMajor = []
        Axios.get('http://localhost:8080/major').then(result => {
            console.log(result.data)
            result.data.forEach(item => {
                dataAllMajor.push(item)
            })
            this.setState({ rows: dataAllMajor })
        })
    }
    render() {
        return (
            <div>
                <center>
                    <h1>Majors</h1>
                    <table border='1' width='50%' cellSpacing="0">
                        <tbody>
                            <tr>
                                <th>No.</th><th>Major</th>
                            </tr>
                            {this.state.rows.map(item => (
                                <tr key={item.major_id}>
                                    <td>{item.major_id}</td>
                                    <td><Link to={`/sbymajor/${item.major_id}`}>{item.major_name}</Link></td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </center>
            </div>
        )
    }
}

export default Major
