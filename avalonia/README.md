```sh
# creation
cd ~/code/PeriodicTasks/avalonia
sudo snap install dotnet-sdk-90 --classic # installs to /snap/dotnet-sdk-90/current/usr/bin/dotnet
export DOTNET_ROOT=/snap/dotnet-sdk-90/current

echo 'alias dotnet="/snap/dotnet-sdk-90/current/usr/bin/dotnet"' >> ~/.bashrc
source ~/.bashrc

dotnet new install Avalonia.Templates
dotnet new avalonia.app -o MyApp
cd MyApp
mv * ..
cd ..
dotnet run

# run
cd ~/code/PeriodicTasks/avalonia
dotnet run

# publish
cd ~/code/PeriodicTasks/avalonia
dotnet publish -c Release -r linux-x64 --self-contained true /p:PublishSingleFile=true
```