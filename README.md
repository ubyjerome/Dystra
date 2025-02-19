### **Dystra – Automated File Sorting Made Simple**  
[![Go](https://img.shields.io/badge/Go-1.2x-blue)](https://go.dev/)  
Dystra is a lightweight, configurable file organizer written in Go. It scans specified directories, identifies files based on their extensions, and automatically sorts them into predefined folders.  

## **Features**  
✅ Automated folder scanning and sorting  
✅ Customizable rules for file organization  
✅ Configurable execution frequency  
✅ Creates missing directories automatically  

## **Installation**  

1. **Download the binary** 
2. **Move it to your home directory**  
   ```sh
   mv dystra ~/  
   chmod +x ~/dystra  
   ```  
3. **Create a `.dystra.config.json` file** in your home directory:  
   ```json
   {
      "frequency": "1min",
      "scan": [
         "~/Downloads",
         "~/Projects"
      ]
   }
   ```  
4. **Create `.dystra.rules.config`** in any directory you want to be sorted:  
   ```
   IF "pdf" THEN "./document"
   IF "mp4" THEN "./videos"
   ```  

## **Usage**  
Run Dystra manually:  
```sh
~/dystra
```  
Or add it to a cron job for automated execution.  

## **How It Works**  
- **Dystra scans folders** listed in `.dystra.config.json`  
- **It checks for `.dystra.rules.config`** in each directory  
- **It moves files** based on extension rules  
- **Missing folders are automatically created**  

## **Contributing**  
Contributions are welcome! Feel free to fork the repo and submit a PR.  

## **License**  
This project is licensed under the MIT License.  

## **Author**  
📌 Developed by **Ubongabasi Jerome**  
🔗 [GitHub](https://github.com/ubyjerome/Dystra)  

