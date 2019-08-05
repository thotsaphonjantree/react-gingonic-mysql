import React, { Component } from 'react'
import Axios from 'axios'

export class major extends Component {
    constructor(props) {
        super(props)

        this.state = { rows: [] }
    }

    componentDidMount() {

        { this.myfunction() }
    }
    myfunction = () => {
        var dataAllMajor = []
        Axios.get('http://localhost:8080/majors').then(result => {
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
                    <table border='1' width='50%'>
                        <tbody>
                            <tr>
                                <th>No.</th><th>Major</th>
                            </tr>
                            {this.state.rows.map(item => (
                                <tr key={item.major_id}>
                                    <td>{item.major_id}</td><td>{item.major_name}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </center>
            </div>
        )
    }
}

export default major
