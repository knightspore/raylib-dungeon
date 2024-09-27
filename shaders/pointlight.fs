#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D tex;
uniform sampler2D u_normal;
uniform vec2 u_resolution;
uniform vec2 u_lightPos;
uniform vec3 u_lightColor;
uniform float u_ambient;

void main() {
    // Setup
    vec3 lightColor = vec3(u_lightColor.r / 255.0, u_lightColor.g / 255.0, u_lightColor.b / 255.0);
    vec4 textureColor = texture(tex, fragTexCoord);
    vec4 normal = texture(u_normal, fragTexCoord);
    float specularColor = pow(dot(textureColor.rgb, vec3(0.299, 0.587, 0.114)), 2.0);

    // Ambient Light
    vec3 lighting = textureColor.rgb * u_ambient;

    // Current Light Direction
    vec2 pos = vec2(u_lightPos.x / u_resolution.x, 1.0 - u_lightPos.y / u_resolution.y);
    vec2 lightDir2d = pos - fragTexCoord;
    vec2 lightDirNorm = normalize(lightDir2d);
    vec3 lightDir3 = normalize(vec3(lightDirNorm, 200.0));
    float distance = length(lightDir2d);

    // Diffuse Lighting
    float diffuseStrength = max(dot(normal.rgb, lightDir3), 0.0);
    float attenuation = 1.0 / (1.0 + 5.0 * distance + 4.0 * distance * distance);
    vec3 diffuse = diffuseStrength * attenuation * textureColor.rgb * lightColor;
    vec3 specular = specularColor * attenuation * lightColor;

    finalColor = vec4(lighting + diffuse + specular, textureColor.a * fragColor.a);
}
