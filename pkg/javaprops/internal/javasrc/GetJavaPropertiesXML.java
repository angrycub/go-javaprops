package io.nomadproject.java;

import java.io.IOException;
import java.util.Properties;

/**
* Returns JSON representation of System.properties().
*
* This application takes no arguments and writes the JSON to STDOUT.
*/
public class GetJavaPropertiesXML
{
    public static void main( String[] args )
    {
        Properties props = System.getProperties();
        try {
            props.storeToXML( System.out, "system.properties", "UTF-8");
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    };
}
