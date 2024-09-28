#version 410 core

in vec2 fragTexCoord;
in vec4 fragColor;
out vec4 finalColor;
uniform sampler2D tex;

uniform int size = 2; // blur size
uniform float separation = 3.0; // spread
uniform float threshold = 0.8;
uniform float amount = 1.0;

uniform float chromaticAberration = 0.005;

void main() {
    vec2 texSize = textureSize(tex, 0).xy;
    float dist = length(fragTexCoord.y - 0.5);
    
    vec4 resultR = vec4(0.0);
    vec4 resultG = vec4(0.0);
    vec4 resultB = vec4(0.0);
    float count = 0.0;
    
    for (int x = -size; x <= size; x++) {
        for (int y = -size; y <= size; y++) {
            vec2 offset = (vec2(x, y) * separation * dist) / texSize;
            
            // Sample each color channel with a slight offset
            vec4 colorR = texture(tex, fragTexCoord + offset + vec2(chromaticAberration, 0.0) * dist * 2.0);
            vec4 colorG = texture(tex, fragTexCoord + offset);
            vec4 colorB = texture(tex, fragTexCoord + offset - vec2(chromaticAberration, 0.0) * dist * 2.0);
            
            // Apply threshold
            float valueR = max(colorR.r, max(colorR.g, colorR.b));
            float valueG = max(colorG.r, max(colorG.g, colorG.b));
            float valueB = max(colorB.r, max(colorB.g, colorB.b));
            
            if (valueR > threshold) colorR = vec4(0.0);
            if (valueG > threshold) colorG = vec4(0.0);
            if (valueB > threshold) colorB = vec4(0.0);
            
            resultR += colorR;
            resultG += colorG;
            resultB += colorB;
            count++;
        }
    }
    
    vec4 result = vec4(resultR.r, resultG.g, resultB.b, (resultR.a + resultG.a + resultB.a) / 3.0) / count;
    
    finalColor = fragColor * mix(vec4(1.0), result, amount);
}
