# Build stage
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src
COPY ["Duckploy/Duckploy.csproj", "Duckploy/"]
RUN dotnet restore "Duckploy/Duckploy.csproj"
COPY . .
WORKDIR "/src/Duckploy"
RUN dotnet publish "Duckploy.csproj" -c Release -o /app/publish --no-restore

# Remove NuGet package cache
RUN rm -rf ~/.nuget

# Runtime stage
FROM mcr.microsoft.com/dotnet/runtime:8.0-alpine AS runtime
WORKDIR /app
COPY --from=build /app/publish .
ENTRYPOINT ["dotnet", "Duckploy.dll"]
