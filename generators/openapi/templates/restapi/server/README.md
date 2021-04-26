## How to extend this REST Api

1. Add new entries in api/<%=openApiGenPackage%>.yaml
2. Run make generate-<%=openApiGenPackage%>
3. Add additional implementations for generated rest endpoints in <%=genOutputPath%>
