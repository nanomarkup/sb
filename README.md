# Smart Builder
Smart Builder is the next generation of building applications using a new format of application requirements (.sa) to describe the functionality of the application and independent components (.sc) for implementation the application. It generates an intermediate Smart Builder module (.sb) needed for generating sources of the application. 

Please read "readme.txt" documentation about all public items of the root package.

### Workflow
Tech engineer works with application requrements (.sa file/s).  
Software engineer develops independent components (.sc file/s).  
The Smart Builder chooses the best components for the application (.sb file/s), generates sources, and builds them using a native compiler.

1. Create new application (.sa file/s)
2. Generate intermediate .sb file/s use .sa and .sc files
3. Generate sources of the application
4. Build the application

### Folder Structure
- **.** includes the build scripts
- **app** includes sources of this project
- **bin** includes binary files
- **cmd** includes CLI implementation
- **helper** includes helper packages
- **plugins** includes sources of using plugins (HashiCorp's [Go Plugin System](https://github.com/hashicorp/go-plugin))
- **samples** includes samples for testing the application
