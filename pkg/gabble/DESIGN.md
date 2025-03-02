# Gabble Design

## Package Structure

### Irods

top level shared structure. 

- hold the common configuration and allow overrides
- shared expensive refs
- track the environment (what grids, what versions, discovered capabilities)
- serve as a factory for IrodsContext


### IrodsContext

- wraps a connection between client and server
- account and login information
- connection status
- per connection configuration and tuning
- factory for services
- connection lifecycle




