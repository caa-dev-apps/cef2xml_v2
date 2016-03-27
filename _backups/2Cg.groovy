@Grab(group='joda-time', module='joda-time', version='2.3')
            
import org.joda.time.DateTime
import org.joda.time.DateTimeUtils

def l_drive = "C:", l_root = null, l_tag = null
def scriptDir = getClass().protectionDomain.codeSource.location.path
def matcher = scriptDir =~ /^(.*[\/\\](.*))[\/\\]_backups(.*)$/

if (matcher.matches()) {
    l_root = matcher[0][1]
    l_tag = matcher[0][2]
    
    if (this.args.size() >= 1) {
        def matcher2 = this.args[0] =~ /^((.):)$/
        
        if (matcher2.matches()) {
            l_drive = matcher2[0][1]
        
            if (this.args.size() >= 2) {
                l_tag = this.args[1]
            }    
        }
    }
}

if(l_drive && l_root && l_tag) {
    def m_backupFolder = l_drive + '/_backups'
    def m_sourceFolder = l_root

    def m_now = new DateTime()
                          
    def dest = l_drive + 
               '/_backups' + 
                "/" +
                m_now.toString("YYYY.MM.MMMM") +
                "/" +
                m_now.toString("YYYY.MM.dd") +
                "/" +
                (m_now =~ /[:-]/).replaceAll("") +
                "__" +
                l_tag +
                ".zip"
                          
    println("")
    println("-- src:  " + m_sourceFolder)
                          
    new AntBuilder().zip(destfile: dest,  
        basedir: m_sourceFolder,  
        includes: "**/*.*", 
        compress: true,
        excludes: "**/*.doc")     
}
else {
    print('Please check command line arguments')
    print('No args = default => C: and Parent Folder as Tag')
    print('else')
    print('arg[1] = Drive e.g C: or M:')
    print('arg[2] = optional tag string')
}



