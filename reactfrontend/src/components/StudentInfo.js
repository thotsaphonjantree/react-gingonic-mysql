import React, { Component } from 'react'
import Axios from 'axios'
export default class StudentInfo extends Component {
    constructor(props) {
        super(props)
        this.state = { student: '', major: [], rows: [''] }

        this.handleDelete = this.handleDelete.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
    }
    componentDidMount() {
        { this.getStudentbyID() }
        { this.getMajor() }
    }

    getMajor = () => {
        var dataAllMajor = []
        Axios.get('http://localhost:8080/majors').then(result => {
            console.log(result.data)
            result.data.forEach(item => {
                dataAllMajor.push(item)
            })
            this.setState({ rows: dataAllMajor })
        })
    }

    getStudentbyID() {
        const { match: { params } } = this.props;
        Axios.get(`http://localhost:8080/student/${params.studentId}`).then(result => {
            this.setState({ student: result.data, major: result.data.Major })
            console.log(result.data);
            console.log(this.state.student.Major);
        })
    }

    handleDelete() {
        const { match: { params } } = this.props;
        Axios.delete(`http://localhost:8080/student/${params.studentId}`)
            .then(() => {
                console.log('user deleted');
            });
        window.location = '/students/';
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = event.target.name;
        this.setState(prevState => ({
            ...prevState, student: {
                ...prevState.student, [name]: value
            }
        }));
    }

    handleSubmit = event => {
        event.preventDefault();
        const { match: { params } } = this.props;
        Axios.put(`http://localhost:8080/student/${params.studentId}`, {
            student_code: this.state.student.student_code,
            first_name: this.state.student.first_name,
            last_name: this.state.student.last_name,
            major_id: parseInt(this.state.student.major_id)
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
            <div>
                <center>
                    <h1>Edit Student</h1>
                    <form onSubmit={this.handleSubmit}>
                        <label>
                            Student ID:
                <input type="text" name="student_code" value={this.state.student.student_code || ''} onChange={this.handleInputChange} />
                        </label><br />
                        <label>
                            First Name:
                <input type="text" name="first_name" value={this.state.student.first_name || ''} onChange={this.handleInputChange} />
                        </label><br />
                        <label>
                            Last Name:
                <input type="text" name="last_name" value={this.state.student.last_name || ''} onChange={this.handleInputChange} />
                        </label><br />
                        <label>
                            Major:
                        <select name="major_id" value={this.state.student.major_id} onChange={this.handleInputChange}>
                                <option value={this.state.major.major_id} disabled>{this.state.major.major_name}</option>
                                {this.state.rows.map((item, index) => (
                                    <option value={item.major_id} key={index} >{item.major_name}</option>
                                ))}
                            </select>
                        </label><br />
                        <button type="submit">Edit</button>
                        <button onClick={this.handleDelete}>
                            Delete
                </button>
                    </form>
                </center>


            </div>
        )
    }
}
