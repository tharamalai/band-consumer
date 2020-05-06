import React from 'react'
import './App.css'

import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
} from 'react-router-dom'

import theme from './ui/theme'
import { ThemeProvider } from 'styled-components'
import { CosmosHubProvider } from 'contexts/CosmosHub'
import { MeichainProvider } from 'contexts/MeichainContext'

import LandingPage from 'pages/Landing'

export default () => (
  <ThemeProvider theme={theme}>
    <CosmosHubProvider>
      <MeichainProvider>
        <Router>
          <Switch>
            <Route exact path="/" component={LandingPage} />
            <Route path="/" render={() => <Redirect to="/" />} />
          </Switch>
        </Router>
      </MeichainProvider>
    </CosmosHubProvider>
  </ThemeProvider>
)
