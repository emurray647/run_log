import logo from './logo.svg';
import './App.css';

import { Switch, Route, Link } from "react-router-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import Navbar from "./components/Navbar";
import ActivityList from "./components/ActivityList"
import Activity from "./components/Activity"

function App() {
    return (
        <div className="App">
            <Navbar />

            <div className="container mt-3">
                <Switch>
                    <Route exact path={["/activities"]} component={ActivityList} />
                    <Route exact path={["/activities/:id"]} component={Activity} />
                    {/* <Route 
                        path="/restaurants/:id/review"
                        render={(props) => (
                        <AddReview {...props} user={user} />
                        )}
                    />
                    <Route 
                        path="/restaurants/:id"
                        render={(props) => (
                        <Restaurant {...props} user={user} />
                        )}
                    />
                    <Route 
                        path="/login"
                        render={(props) => (
                        <Login {...props} login={login} />
                        )}
                    /> */}
                </Switch>
            </div>  
        </div>
    );
}

export default App;
